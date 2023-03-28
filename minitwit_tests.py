# minitwit_tests.py

import unittest

class TestStringMethods(unittest.TestCase):

    def test_will_always_be_true(self):
        self.assertEqual(True, True)

    def test_isupper(self):
        self.assertTrue('HELLO'.isupper())
        self.assertFalse('Hello'.isupper())

    def test_split(self):
        s = 'hello world'
        self.assertEqual(s.split(), ['hello', 'world'])
        # check that s.split fails when the separator is not a string


if __name__ == '__main__':
    unittest.main()