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
