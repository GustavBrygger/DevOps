# minitwit_tests.py

import unittest
import psycopg2
import os

BASE_URL = 'http://127.0.0.1:5432'

DATABASE = "postgres"
USERNAME = 'postgres'
PWD = 'postgres'
PORT = '5432'


def init_db():

    try:
        """Connect database ."""
        conn = psycopg2.connect(
            host="db",
            database=DATABASE,
            user=USERNAME,
            password=PWD,
            port=PORT)
        
        cur = conn.cursor()

                
        # execute a statement
        print('PostgreSQL database version:')
        cur.execute('SELECT version()')

        # Run the SQL file to the database
        cur.execute(open("schema.sql", "r").read())

        # display the PostgreSQL database server version
        db_version = cur.fetchone()
        print(db_version)
        
        # close the communication with the PostgreSQL
        cur.close()
    except (Exception, psycopg2.DatabaseError) as error:
        print(error)
    finally:
        if conn is not None:
            conn.close()
            print('Database connection closed.')

    
    




class TestStringMethods(unittest.TestCase):


    


    def test_will_always_be_true(self):
        self.assertEqual(True, True)

    def test_will_always_be_false(self):
        self.assertEqual(True, False)

    def test_isupper(self):
        self.assertTrue('HELLO'.isupper())
        self.assertFalse('Hello'.isupper())

    def test_split(self):
        s = 'hello world'
        self.assertEqual(s.split(), ['hello', 'world'])
        # check that s.split fails when the separator is not a string


if __name__ == '__main__':
    init_db()

    unittest.main()