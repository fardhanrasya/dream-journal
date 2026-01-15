$ErrorActionPreference = "Stop"
[Net.ServicePointManager]::SecurityProtocol = [Net.SecurityProtocolType]::Tls12

$Repo = "fardhanrasya/dream-journal"
$AppName = "dream"
$InstallDir = "$env:USERPROFILE\.dream-journal"
$BinDir = "$InstallDir\bin"

# Determine Architecture
$Arch = if ($env:PROCESSOR_ARCHITECTURE -eq "AMD64") { "x86_64" } else { "arm64" } # Simplification, mostly x64 on Windows

try {
    # Get latest release info
    Write-Host "Fetching latest release information..."
    $ReleasesUrl = "https://api.github.com/repos/$Repo/releases/latest"
    $Release = Invoke-RestMethod -Uri $ReleasesUrl
    $Version = $Release.tag_name
    Write-Host "Latest version: $Version"

    $AssetPattern = "*Windows_$Arch.zip"
    $Asset = $Release.assets | Where-Object { $_.name -like $AssetPattern } | Select-Object -First 1

    if (-not $Asset) {
        throw "Could not find a release asset for Windows $Arch"
    }
    $DownloadUrl = $Asset.browser_download_url
}
catch {
    Write-Warning "Unable to connect to GitHub API or find asset. Attempting direct download fallback."
    $Version = "latest"
    $DownloadUrl = "https://github.com/$Repo/releases/latest/download/dream-journal_Windows_$Arch.zip"
}

# Download
$ZipPath = "$env:TEMP\dream_journal.zip"
Write-Host "Downloading $DownloadUrl..."
Invoke-WebRequest -Uri $DownloadUrl -OutFile $ZipPath

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
