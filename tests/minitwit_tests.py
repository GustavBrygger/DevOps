# minitwit_tests.py

import unittest
import requests
import psycopg2
import os

BASE_URL = 'http://127.0.0.1:5432'

PORT = '5432'


def init_db():

    try:
        """Connect database ."""
        conn = psycopg2.connect(
            host="db",
            database='postgres',
            user='postgres',
            password='postgres',
            port=PORT)
        
        cur = conn.cursor()

                
        # display the PostgreSQL database server version
        print('PostgreSQL database version:')
        cur.execute('SELECT version()')
        db_version = cur.fetchone()
        conn.commit()
        print(db_version)

        #cur.execute('SELECT * FROM users')
        #answer = cur.fetchone()
        #print(answer)
        #conn.commit()
        
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

    def test_if_register_works(self):
        API_ENDPOINT = "http://server:8080/register"
        data = {'email':'gustav.brygger',
                'pwd':'password',
                'username':'gubr'}
        
        r = requests.post(url = API_ENDPOINT, data = data)
        self.assertEqual(r.status_code, 204)

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
