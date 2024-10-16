import sqlite3
import logging
import time
import os
from sample.color.printColor import printBlackOnWhite, printGreen, printRed


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
    if not os.path.isfile("mypasswords.db"):
        printBlackOnWhite("Creating database...")
        conn = None
        cur = None
        try:
            conn = sqlite3.connect("mypasswords.db")
            cur = conn.cursor()

            for statment in sql_statments:
                cur.execute(statment)
            return True
        except sqlite3.Error as e:
            logging.error(e)
        finally:
            if cur is not None:
                cur.close()
            if conn is not None:
                conn.close()
    else:
        printGreen("Database already exists")
        return True

    return False


def dblogin(user: str, password: str) -> bool:
    sql = "SELECT password from users where login=?"
    conn = None
    cur = None
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
        if cur is not None:
            cur.close()
        if conn is not None:
            conn.close()

    return False


def dbsignUp(user: str, password: str) -> bool:
    sql = "INSERT INTO users (login, password, lastlogin) VALUES (?, ?, ?)"
    conn = None
    cur = None
    try:
        conn = sqlite3.connect("mypasswords.db")
        cur = conn.cursor()

        cur.execute(sql, [user, password, time.time()])
        conn.commit()
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        if cur is not None:
            cur.close()
        if conn is not None:
            conn.close()

    return True


def dbgetPassword(passName: str, username: str) -> list:
    sql = "SELECT name, password, url, folder_id FROM passwords WHERE user_id=? AND name=?"
    conn = None
    cur = None
    try:
        conn = sqlite3.connect("mypasswords.db")
        conn.row_factory = row_to_dict
        cur = conn.cursor()

        cur.execute(sql, (username, passName))
        return [cur.fetchone()]
    except sqlite3.Error as e:
        logging.error(e)
    finally:
        if cur is not None:
            cur.close()
        if conn is not None:
            conn.close()

    return []


def dbgetAllPasswords(username: str) -> list:
    sql = "SELECT name, password, url FROM passwords WHERE user_id=?"
    passDict = []
    conn = None
    cur = None
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
        if cur is not None:
            cur.close()
        if conn is not None:
            conn.close()

    return passDict


def dbinsertPassword(passName: str, password: str, url: str, folder: str, username: str) -> bool:
    sql = "INSERT INTO passwords (name, password, url, folder_id, user_id) VALUES(?, ?, ?, ?, ?)"
    conn = None
    cur = None
    try:
        conn = sqlite3.connect("mypasswords.db")
        cur = conn.cursor()

        cur.execute(sql, (passName, password, url, folder, username))
        conn.commit()
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        if cur is not None:
            cur.close()
        if conn is not None:
            conn.close()

    return True


def dbDeletePassword(passName: str, username: str) -> bool:
    sql = "DELTE FROM passwords WHERE name=? and user_id=?"
    conn = None
    cur = None
    try:
        conn = sqlite3.connect("mypasswords.db")
        cur = conn.cursor()

        cur.execute(sql, (passName, username))
        conn.commit()
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        if cur is not None:
            cur.close()
        if conn is not None:
            conn.close()

    return True

def dbChangePasswordName(passName: str, username: str, **kwargs) -> bool:
    set_clause = ', '.join([f"{key} = ?" for key in kwargs.keys()])

    sql = f"UPDATE passwords SET {set_clause} WHERE name=? and user_id=?"
    conn = None
    cur = None
    try:
        conn = sqlite3.connect("mypasswords.db")
        cur = conn.cursor()

        cur.execute(sql, (passName, username))
        conn.commit()
    except sqlite3.Error as e:
        logging.error(e)
        return False
    finally:
        if cur is not None:
            cur.close()
        if conn is not None:
            conn.close()

    return True

def row_to_dict(cursor: sqlite3.Cursor, row: sqlite3.Row) -> dict:
    data = {}
    for idx, col in enumerate(cursor.description):
        data[col[0]] = row[idx]
    return data
