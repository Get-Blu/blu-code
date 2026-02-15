const fs = require('fs');
const https = require('https');
const path = require('path');
const os = require('os');
const { version } = require('../package.json');

const platform = process.platform;
const arch = process.arch;

// Map OS
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
    console.error(`Unsupported OS: ${platform}`);
    process.exit(1);
}

// Map Architecture
let goarch = '';
if (arch === 'x64') {
    goarch = 'amd64';
} else if (arch === 'arm64') {
    goarch = 'arm64';
} else {
    console.error(`Unsupported architecture: ${arch}`);
    process.exit(1);
}

// Construct URL
// Artifact names from GitHub Release: blu-code-linux-amd64, blu-code-windows-amd64.exe
const binName = `blu-code-${goos}-${goarch}${ext}`;
const url = `https://github.com/Get-Blu/blu-code/releases/download/v${version}/${binName}`;

console.log(`Downloading ${url} ...`);

const download = (url, dest) => {
    return new Promise((resolve, reject) => {
        const file = fs.createWriteStream(dest);
        https.get(url, (response) => {
            if (response.statusCode === 302 || response.statusCode === 301) {
                file.close();
                fs.unlinkSync(dest); // Clean up partial file before redirect
                return resolve(download(response.headers.location, dest));
            }

            if (response.statusCode !== 200) {
                file.close();
                fs.unlinkSync(dest);
                return reject(new Error(`Failed to download binary: HTTP ${response.statusCode}`));
            }

            response.pipe(file);
            file.on('finish', () => {
                file.close();
                const stats = fs.statSync(dest);
                if (stats.size === 0) {
                    fs.unlinkSync(dest);
                    return reject(new Error('Downloaded binary is empty'));
                }
                resolve();
            });
        }).on('error', (err) => {
            file.close();
            if (fs.existsSync(dest)) fs.unlinkSync(dest);
            reject(err);
        });
    });
};

const destDir = path.join(__dirname, '..', 'bin');
if (!fs.existsSync(destDir)) {
    fs.mkdirSync(destDir, { recursive: true });
}

const destFile = path.join(destDir, `blu${ext}`);

download(url, destFile)
    .then(() => {
        console.log(`Download complete: ${destFile}`);
        if (platform !== 'win32') {
            fs.chmodSync(destFile, 0o755);
        }
    })
    .catch((err) => {
        console.error(`Error: ${err.message}`);
        process.exit(1);
    });
