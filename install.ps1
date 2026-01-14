$ErrorActionPreference = "Stop"

$Repo = "fardhanrasya/dream-journal"
$AppName = "dream"
$InstallDir = "$env:USERPROFILE\.dream-journal"
$BinDir = "$InstallDir\bin"

# Get latest release info
Write-Host "Fetching latest release information..."
$ReleasesUrl = "https://api.github.com/repos/$Repo/releases/latest"
$Release = Invoke-RestMethod -Uri $ReleasesUrl

$Version = $Release.tag_name
Write-Host "Latest version: $Version"

# Determine Architecture
$Arch = if ($env:PROCESSOR_ARCHITECTURE -eq "AMD64") { "x86_64" } else { "arm64" } # Simplification, mostly x64 on Windows
$AssetPattern = "*Windows_$Arch.zip"

# Find asset
$Asset = $Release.assets | Where-Object { $_.name -like $AssetPattern } | Select-Object -First 1

if (-not $Asset) {
    Write-Error "Could not find a release asset for Windows $Arch"
}

# Download
$ZipPath = "$env:TEMP\dream_journal.zip"
Write-Host "Downloading $($Asset.browser_download_url)..."
Invoke-WebRequest -Uri $Asset.browser_download_url -OutFile $ZipPath

# Install
Write-Host "Installing to $BinDir..."
if (-not (Test-Path $BinDir)) { New-Item -Path $BinDir -ItemType Directory -Force | Out-Null }

Expand-Archive -Path $ZipPath -DestinationPath $BinDir -Force
Remove-Item $ZipPath

# Setup PATH
$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$BinDir*") {
    Write-Host "Adding $BinDir to PATH..."
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$BinDir", "User")
    $env:Path += ";$BinDir"
    Write-Host "PATH updated. You may need to restart your terminal."
}

# Setup env var for EDITOR (Default to notepad if not set)
if (-not [Environment]::GetEnvironmentVariable("EDITOR", "User")) {
    Write-Host "Setting EDITOR environment variable to 'notepad'..."
    [Environment]::SetEnvironmentVariable("EDITOR", "notepad", "User")
}

Write-Host "`nSuccess! Dream Journal $Version installed."
Write-Host "Run 'dream' to start."
