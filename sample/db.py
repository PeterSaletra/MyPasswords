from os.path import isfile
import sqlite3
import logging
import time
import os
import json


def createNewDB() -> bool:
    sql_statments = [
        """
            CREATE TABLE IF NOT EXISTS users (
                login text NOT NULL PRIMARY KEY,
                password text NOT NULL,
                lastlogin text NOT NULL
            );
            """,
        """
            CREATE TABLE IF NOT EXISTS folders(
                id INTEGER NOT NULL,
                name text NOT NULL
            );
            """,
        """
            CREATE TABLE IF NOT EXISTS userFolder(
                user_id text NOT NULL,
                folder_id INTEGER NOT NULL,
                FOREIGN KEY (user_id) REFERENCES users (login),
                FOREIGN KEY (folder_id) REFERENCES folders (id)
            );
            """,
        """
            CREATE TABLE IF NOT EXISTS passwords(
                name text NOT NULL,
                password text NOT NULL,
                url text,
                folder_id INTEGER,
                user_id text NOT NULL,
                FOREIGN KEY (folder_id) REFERENCES folders (id),
                FOREIGN KEY (user_id) REFERENCES users (login)
            );
            """,
    ]
    print("Creating database...")
    if not os.path.isfile("mypasswords.db"):
        try:
            conn = sqlite3.connect("mypasswords.db")
            cur = conn.cursor()

            for statment in sql_statments:
                cur.execute(statment)
            return True
        except sqlite3.Error as e:
            logging.error(e)
        finally:
            cur.close()
            conn.close()
    else:
        print("Database already exists. Adding user to existing database")
        return True

    return False


def dblogin(user: str, password: str) -> bool:
    sql = "SELECT password from users where login=?"

    try:
        conn = sqlite3.connect("mypasswords.db")
        conn.row_factory = row_to_dict
        cur = conn.cursor()

        cur.execute(sql, (user,))
        row = cur.fetchone()
        if password == row["password"]:
            return True
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        cur.close()
        conn.close()

    return False


def dbsignUp(user: str, password: str) -> bool:
    sql = "INSERT INTO users (login, password, lastlogin) VALUES (?, ?, ?)"
    try:
        conn = sqlite3.connect("mypasswords.db")
        cur = conn.cursor()

        cur.execute(sql, [user, password, time.time()])
        conn.commit()
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        cur.close()
        conn.close()

    return True


def dbgetPassword(passName: str, username: str) -> list:
    sql = "SELECT name, password, url, folder_id FROM passwords WHERE user_id=? AND name=?"

    try:
        conn = sqlite3.connect("mypasswords.db")
        conn.row_factory = row_to_dict
        cur = conn.cursor()

        cur.execute(sql, (username, passName))
        return [cur.fetchone()]
    except sqlite3.Error as e:
        logging.error(e)
    finally:
        cur.close()
        conn.close()

    return []


def dbgetAllPasswords(username: str) -> list:
    sql = "SELECT name, password, url FROM passwords WHERE user_id=?"
    passDict = []
    try:
        conn = sqlite3.connect("mypasswords.db")
        conn.row_factory = row_to_dict
        cur = conn.cursor()

        cur.execute(sql, [username])
        for row in cur.fetchall():
            passDict.append(row)

    except sqlite3.Error as e:
        logging.error(e)
    finally:
        cur.close()
        conn.close()

    return passDict


def dbinsertPassword(
    passName: str, password: str, url: str, folder: str, username: str
) -> bool:
    sql = "INSERT INTO passwords (name, password, url, folder_id, user_id) VALUES(?, ?, ?, ?, ?)"
    try:
        conn = sqlite3.connect("mypasswords.db")
        cur = conn.cursor()

        cur.execute(sql, (passName, password, url, folder, username))
        conn.commit()
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        cur.close()
        conn.close()

    return True


def dbDeletePassowrd(passName: str, username: str) -> bool:
    sql = ""
    try:
        conn = sqlite3.connect("mypasswords.db")
        cur = conn.cursor()

        cur.execute(sql, (passName, username))
        conn.commit()
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        cur.close()
        conn.close()

    return True


def row_to_dict(cursor: sqlite3.Cursor, row: sqlite3.Row) -> dict:
    data = {}
    for idx, col in enumerate(cursor.description):
        data[col[0]] = row[idx]
    return data
