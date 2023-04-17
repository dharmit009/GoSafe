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


## math/rand vs crypto/rand

This code generates a password with a length of 16 characters by creating a
byte slice of length passwordLength and calling rand.Read() to fill it with
random bytes.

```golang

package main

import (
    "crypto/rand"
    "fmt"
)

func main() {
    const passwordLength = 16

    password := make([]byte, passwordLength)
    _, err := rand.Read(password)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(password))
}

```

The generated password is printed to the console using fmt.Println(). Note that
rand.Read() returns the number of bytes read and an error, which we ignore by
using the blank identifier \(_\). Also, rand.Read() returns a cryptographically
secure random byte slice, which is more suitable for generating passwords than
using the 'math/rand' package.

# GUIDE:

## Define the requirements:

First, define the requirements of your password manager. Determine the features
that you want to include, such as generating strong passwords, securely storing
passwords, syncing passwords across devices, and auditing your passwords.

## Choose a programming language and platform:

Choose a programming language and platform that you're comfortable with and
that can support the features you want to include in your password manager.
Common languages for password managers include Python, Java, and C++. You'll
also need to choose a database to store your passwords securely.

## Design the user interface:

Design the user interface for your password manager. This should be intuitive,
user-friendly, and secure. Consider including features such as password
strength indicators, autofill login credentials, and two-factor authentication.

## Implement password encryption:

Implement strong encryption algorithms to encrypt passwords before storing them
in the database. This is a critical component of your password manager to
ensure that your passwords are secure.

## Develop the password generator:

Develop a password generator that can generate strong and unique passwords for
all your online accounts. This should include options for specifying the
length, complexity, and character types.

## Add syncing and backup capabilities:

Add syncing and backup capabilities to your password manager, so you can access
your passwords from multiple devices and back up your passwords in case of data
loss.

## Test and debug your password manager:

Test your password manager thoroughly to ensure that it's secure, reliable, and
user-friendly. Test for vulnerabilities such as SQL injection, cross-site
scripting, and buffer overflow attacks.

## Release and maintain your password manager:

Release your password manager to the public, and maintain it by releasing
updates and fixing bugs as they arise. You'll also need to stay up to date with
the latest security best practices and patch any vulnerabilities as soon as
possible.
# Creative Commons Attribution-NonCommercial-ShareAlike 4.0 International License.

This work is licensed under the Creative Commons
Attribution-NonCommercial-ShareAlike 4.0 International License. To view a copy
of this license, visit http://creativecommons.org/licenses/by-nc-sa/4.0/ or
send a letter to Creative Commons, PO Box 1866, Mountain View, CA 94042, USA.
# MISC:

**To create a GO Module for your project:**

> go mod init github.com/username/repo-name
> go mod init github.com/dharmit009/gopass
# References:

* [Go Official Documentation](https://go.dev/doc/)
* [Go Package Documentation](https://pkg.go.dev/)
* [Fyne Official Documentation](https://developer.fyne.io/)

## Streamers:

* rwxrob
    * [Twitch](https://twitch.tv/rwxrob)
    * [Youtube](https://youtube.com/rwxrob)
    * [Github](https://www.github.com/rwxrob)
* [ThePrimeAgen](https://youtube.com/theprimeagen)
* [NeuralNine](https://youtube.com/neuralnine)
* [Hitesh Choudary](https://www.youtube.com/@HiteshChoudharydotcom)
* [Computerphile](https://youtube.com/computerphile)

# REQUIREMENTS:

**Q1. What are your requirements for password manager?**

    **Within Scope:**

    1. Add, Remove, Store, & Update passwords.
    1. Automatic Password Generator.
    1. Cross Compiled Password Manager.

    **Out Of Scope:**

    1. 2FA (TOTP) Support.
    1. Password Strength Indicator.
    1. Synchronous Among Various Devices.


**Q2. What Programming Language and Platform are you going to use?**

> Golang

**Q3. Which Library are you going to use for GUI?**

> FyneGUI

**Q4. Which password encryption algorithm are you going to use?**

> Depends on NueralNine Video

**Q5. Which Database are you going to use for storing passwords?**

> maybe just a json file.

**Q6. Which operating systems are your target platforms?**

> Almost all apart from Apple ecosystem.
# TO-DO:

[✅] Setup Golang and FyneGUI

[  ] Create package with name `passutil` which implements the following functions:
    * [✅] generatePass() - Generates and returns a new password.
    * [✅] checkPass() - Used to check the strength of password returns a score out of 10.
    * [] encryptPass() - Used to encrypt the password.
    * [] decryptPass() - used to decrypt the password.

[⚠️ ] Pending Decisions !!!
    * [] Whether to implement Database or not?
        * [] If yes, then look into SQLite3.
        * [] Also check how to synchronize the data between multiple devices.
    * [] Encryption Algorithm?
    * [] Multi-User?

[] Things you need inside your password manager:
    * [] Username and it's Master Password.
    * [] User should be able to save multiple passwords.
    * [] Functionalities like: add, remove, update, logging.
    * [] FyneGUI which implements all the Functionalities.
