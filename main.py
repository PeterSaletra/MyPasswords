from sample.core import main
import argparse

if __name__ == "__main__":
    parser = argparse.ArgumentParser(
        description="MyPasswords is password manager with localy stored passwods.\n\n Keep in my that is educacional project and you are not should use it as your password manager!!"
    )
    parser.add_argument(
        "-l", "--login", help="Use to login as a user. Example: --login username"
    )
    parser.add_argument(
        "-s", "--sign", help="Use to sign up as new user. Example: --sign"
    )
    parser.add_argument(
        "-p", "--passwd", help="Use to enter user password. Example: --passwd password"
    )
    parser.add_argument(
        "-g",
        "--get",
        help="Use to get a specyfic password from your vault. Example: --get facebook",
    )
    parser.add_argument(
        "-a",
        "--add",
        help="Use to add new password to vault as an argument specyfie folder. Example: --add folder",
    )

    args = parser.parse_args()

    if args.login:
        if args.passwd:
            # Update user last login in database
            print(args.login + " " + args.passwd)
            if args.get:
                # Get user password and update user last login in database
                pass
        else:
            # Enter password function
            pass
    elif args.get:
        # Check user last login and get him password
        pass
    elif args.add:
        # Check user last login and add his password
        pass
    else:
        main()
