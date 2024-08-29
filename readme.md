# костыль!! (для vscode)
Удаляет все плагины  которых нет в extensions/recommendations  и устанавливает недостающие  

## установка 
1 скопипастить в .code-workspace file 
```json
"tasks": {
        "version": "2.0.0",
        "tasks": [
            {
                "label": "INSTALL-delete rec-extensions  automaticly ",
                "type": "shell",
                "windows": {
                    "command": "$all_ext=(code --list-extensions );$config_ext= ((Get-Content  -Raw .vscode/*.code-workspace) -replace '\\/\\/.*\\n' -replace '\\n' -replace '\\s' -replace '\\/\\*[^.](.|\\n)*?\\*\\/'  -replace ',}','}'  -replace ',]',']'|ConvertFrom-Json).extensions.recommendations;foreach($ext in $config_ext){  if ($all_ext | ? { $ext -match $_ }){echo 'Already-installed:$ext'}else{echo 'Install:$ext';code --install-extension $ext}   }$all_ext=';$config_ext=';foreach ($ext in  $all_ext) {if ($config_ext | ? { $ext -match $_ }){echo 'exists in config:$ext '}else{echo 'delete:$ext';code  --uninstall-extension $ext}};$all_ext=(code --list-extensions );foreach ($ext in  $all_ext) {if ($config_ext | ? { $ext -match $_ }){echo 'exists in config:$ext '}else{echo 'delete:$ext';code  --uninstall-extension $ext}};"
                },
                "runOptions": {
                    "runOn": "folderOpen"
                },
                "presentation": {
                    "reveal": "silent",
                    "close": true
                }
            }
        ]
    }

```


2 добавить плагины  в .code-workspace file 

```json

"extensions": {
    "recommendations": [
          // Your extentions 
    ]
}
```
