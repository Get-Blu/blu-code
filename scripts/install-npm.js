#!/usr/bin/env node
'use strict';

const fs = require('fs');
const https = require('https');
const path = require('path');
const os = require('os');
const { execSync } = require('child_process');

// Try to get version from package.json
let version;
try {
    const pkg = require('../package.json');
    version = pkg.version;
} catch (e) {
    console.error('Could not read package.json:', e.message);
    process.exit(1);
}

const platform = process.platform;
const arch = process.arch;

// Map Node OS to GOOS
let goos = '';
let ext = '';
if (platform === 'win32') {
    goos = 'windows';
    ext = '.exe';
} else if (platform === 'darwin') {
    goos = 'darwin';
} else if (platform === 'linux') {
    goos = 'linux';
} else {
    console.error(`Unsupported platform: ${platform}`);
    console.error('Please download the binary manually from:');
    console.error(`  https://github.com/Get-Blu/blu-code/releases`);
    process.exit(1);
}

// Map Node arch to GOARCH
let goarch = '';
if (arch === 'x64') {
    goarch = 'amd64';
} else if (arch === 'arm64') {
    goarch = 'arm64';
} else {
    console.error(`Unsupported architecture: ${arch}`);
    console.error('Please download the binary manually from:');
    console.error(`  https://github.com/Get-Blu/blu-code/releases`);
    process.exit(1);
}

// Binary name matches GoReleaser output: blu-linux-amd64, blu-windows-amd64.exe, etc.
const binName = `blu-${goos}-${goarch}${ext}`;
const downloadUrl = `https://github.com/Get-Blu/blu-code/releases/download/v${version}/${binName}`;

const destDir = path.join(__dirname, '..', 'bin');
if (!fs.existsSync(destDir)) {
    fs.mkdirSync(destDir, { recursive: true });
}

const destFile = path.join(destDir, `blu${ext}`);

console.log(`\nDownloading Blu CLI v${version} for ${goos}/${goarch}...`);
console.log(`URL: ${downloadUrl}\n`);

/**
 * Downloads a URL to a local file, following redirects.
 */
function download(url, dest) {
    return new Promise((resolve, reject) => {
        const file = fs.createWriteStream(dest);

        function get(currentUrl) {
            https.get(currentUrl, (response) => {
                // Follow redirects (GitHub releases do a 302 -> S3)
                if (response.statusCode === 301 || response.statusCode === 302) {
                    file.close();
                    if (fs.existsSync(dest)) fs.unlinkSync(dest);
                    const newFile = fs.createWriteStream(dest);
                    file.destroy();
                    return get(response.headers.location);
                }

                if (response.statusCode !== 200) {
                    file.destroy();
                    if (fs.existsSync(dest)) fs.unlinkSync(dest);
                    return reject(
                        new Error(
                            `HTTP ${response.statusCode} — binary not found.\n` +
                            `Make sure release v${version} exists at:\n` +
                            `  https://github.com/Get-Blu/blu-code/releases/tag/v${version}`
                        )
                    );
                }

                response.pipe(fs.createWriteStream(dest));
                response.on('end', () => {
                    const stats = fs.statSync(dest);
                    if (stats.size === 0) {
                        fs.unlinkSync(dest);
                        return reject(new Error('Downloaded binary is empty — the release asset may be missing.'));
                    }
                    resolve();
                });
                response.on('error', reject);
            }).on('error', reject);
        }

        get(url);
    });
}

download(downloadUrl, destFile)
    .then(() => {
        // Make executable on Unix
        if (platform !== 'win32') {
            fs.chmodSync(destFile, 0o755);
        }
        console.log(`✅  Blu CLI installed to: ${destFile}`);
        console.log('    Run "blu --version" to verify the installation.\n');
    })
    .catch((err) => {
        console.error(`\n❌  Installation failed: ${err.message}\n`);
        console.error('Manual install options:');
        console.error('  • curl/wget: https://github.com/Get-Blu/blu-code/releases');
        console.error(`  • Homebrew (macOS/Linux): brew install Get-Blu/tap/blu`);
        process.exit(1);
    });
