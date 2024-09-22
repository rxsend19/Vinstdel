$all_ext = (code --list-extensions );
$config_ext = ((Get-Content  -Raw .vscode/*.code-workspace) -replace '\/\/.*\n' -replace '\n' -replace '\s' -replace '\/\*[^.](.|\n)*?\*\/' -replace ',}', '}' -replace ',]', ']' | ConvertFrom-Json).extensions.recommendations; 
$problem_arr = @();
foreach ($ext in $config_ext) { 
    if ($all_ext | Where-Object { $ext -match $_ }) {
        Write-Host 'Already-installed:$ext'; 
    }
    else {
        Write-Host 'Install:$ext'; code --install-extension $ext;
    } 
};
$all_ext = (code --list-extensions );

foreach ($ext in  $all_ext) {
    if ($config_ext | Where-Object { $ext -match $_ }) {
        Write-Host 'exists in config:$ext'; 
    }
    else { 
        Write-Host 'delete:$ext'; 
        $result = @(code  --uninstall-extension $ext 2>&1);
        if ("Cannot" | Where-Object { $result -match $_ }) {
            $problem_arr += $ext;
        }
    } 
};
foreach ($ext in  $problem_arr) {
    Write-Host 'Re-delete:$ext'; 
    code  --uninstall-extension $ext;
}


# TODO : автоматически выделять аддоны в extensions.json ( если те имеются в code.workspace )