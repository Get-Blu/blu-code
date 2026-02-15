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

const destDir = path.join(__dirname, '..', 'bin');
if (!fs.existsSync(destDir)) {
    fs.mkdirSync(destDir, { recursive: true });
}

const destFile = path.join(destDir, `blu${ext}`);

const file = fs.createWriteStream(destFile);

https.get(url, (response) => {
    if (response.statusCode === 302 || response.statusCode === 301) {
        // Handle redirect
        https.get(response.headers.location, (response) => {
            if (response.statusCode !== 200) {
                console.error(`Failed to download binary: ${response.statusCode}`);
                process.exit(1);
            }
            response.pipe(file);
            file.on('finish', () => {
                file.close();
                console.log(`Download complete: ${destFile}`);
                if (platform !== 'win32') {
                    fs.chmodSync(destFile, 0o755);
                }
            });
        });
    } else if (response.statusCode !== 200) {
        console.error(`Failed to download binary: ${response.statusCode}`);
        process.exit(1);
    } else {
        response.pipe(file);
        file.on('finish', () => {
            file.close();
            console.log(`Download complete: ${destFile}`);
            if (platform !== 'win32') {
                fs.chmodSync(destFile, 0o755);
            }
        });
    }
}).on('error', (err) => {
    fs.unlink(destFile, () => { }); // Delete the file async. (But we don't check result)
    console.error(`Error downloading binary: ${err.message}`);
    process.exit(1);
});
