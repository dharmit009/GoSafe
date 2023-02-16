# Dependencies:

## Windows:

> Download Go from the download page and follow instructions
> Install one of the available C compilers for windows, the following are tested with Go and Fyne:
> MSYS2 with MingW-w64 - msys2.org
> TDM-GCC - tdm-gcc.tdragon.net
> Cygwin - cygwin.com
> In Windows your graphics driver will already be installed, but it is recommended to ensure they are up to date.
> The steps for installing with MSYS2 (recommended) are as follows:
>
> Install MSYS2 from msys2.org
> Once installed do not use the MSYS terminal that opens
> Open “MSYS2 MinGW 64-bit” from the start menu
> Execute the following commands (if asked for install options be sure to choose “all”):
>
>   $ pacman -Syu
>   $ pacman -S git mingw-w64-x86_64-toolchain
> You will need to add /c/Program\ Files/Go/bin and ~/Go/bin to your PATH, for MSYS2 you can paste the following command into your terminal:
>
>   $ echo "export PATH=\$PATH:/c/Program\ Files/Go/bin:~/Go/bin" >> ~/.bashrc

## Debian / Ubuntu:

> sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev

## Fedora:

> sudo dnf install golang gcc libXcursor-devel libXrandr-devel mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel

## Arch Linux:

> sudo pacman -S go xorg-server-devel libxcursor libxrandr libxinerama libxi

## Solus:

> sudo eopkg it -c system.devel golang mesalib-devel libxrandr-devel libxcursor-devel libxi-devel libxinerama-devel

## openSUSE:

> sudo zypper install go gcc libXcursor-devel libXrandr-devel Mesa-libGL-devel libXi-devel libXinerama-devel libXxf86vm-devel

## Void Linux:

sudo xbps-install -S go base-devel xorg-server-devel libXrandr-devel libXcursor-devel libXinerama-devel

## Alpine Linux:

sudo apk add go gcc libxcursor-dev libxrandr-dev libxinerama-dev libxi-dev linux-headers mesa-dev

