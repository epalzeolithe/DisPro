(cls), ($Env:GoOS="linux"), ($Env:GoArch="amd64"), (Go.exe clean -cache -testcache -modcache), (Del ".\resource.syso", ".\Lin_64\DisPro.bin"), (Syso.exe -c=".\resource.json" -o=".\resource.syso"), (Attrib.exe -r +a -i ".\resource.syso" /d /l), (Compact.exe /c /a /i ".\resource.syso"), (Go.exe build -o=".\Lin_64\DisPro.bin" -i -p=4 -x -buildmode=pie -compiler=gc -gcflags="-dwarf=false -e=true -l=false -pack=true -w=true -wb=false" -ldflags="-aslr=true -benchmark=cpu -compressdwarf=true -d=false -extld=gcc -f=false -g=false -race=false -s=true -strictdups=1 -w=true" .\main.go .\socks.go .\linux.go .\const.go), (Attrib.exe -r +a -i ".\Lin_64\DisPro.bin" /d /l), (Compact.exe /c /a /i ".\Lin_64\DisPro.bin")
(cls), ($Env:GoOS="darwin"), ($Env:GoArch="amd64"), (Go.exe clean -cache -testcache -modcache), (Del ".\resource.syso", ".\Dar_64\DisPro.bin"), (Syso.exe -c=".\resource.json" -o=".\resource.syso"), (Attrib.exe -r +a -i ".\resource.syso" /d /l), (Compact.exe /c /a /i ".\resource.syso"), (Go.exe build -o=".\Dar_64\DisPro.bin" -i -p=4 -x -buildmode=pie -compiler=gc -gcflags="-dwarf=false -e=true -l=false -pack=true -race=false -spectre=all -w=true -wb=false" -ldflags="-aslr=true -benchmark=cpu -compressdwarf=true -d=false -extld=gcc -f=false -g=false -race=false -s=true -strictdups=1 -w=true" .\main.go .\socks.go .\darwin.go .\const.go), (Attrib.exe -r +a -i ".\Dar_64\DisPro.bin" /d /l), (Compact.exe /c /a /i ".\Dar_64\DisPro.bin")
(cls), ($Env:GoOS="windows"), ($Env:GoArch="amd64"), (Go.exe clean -cache -testcache -modcache), (Del ".\resource.syso", ".\Win_64\DisPro.bin"), (Syso.exe -c=".\resource.json" -o=".\resource.syso"), (Attrib.exe -r +a -i ".\resource.syso" /d /l), (Compact.exe /c /a /i ".\resource.syso"), (Go.exe build -o=".\Win_64\DisPro.bin" -i -p=4 -x -buildmode=pie -compiler=gc -gcflags="-c=4 -dwarf=false -e=true -l=false -pack=true -race=false -spectre=all -w=true -wb=false" -ldflags="-aslr=true -benchmark=cpu -compressdwarf=true -d=false -extld=gcc -f=false -g=false -race=false -s=true -strictdups=1 -w=true" .\main.go .\socks.go .\windows.go .\const.go), (Attrib.exe -r +a -i ".\Win_64\DisPro.bin" /d /l), (Compact.exe /c /a /i ".\Win_64\DisPro.bin")
