import unittest
from anagrampkg.anagramlib import apply_anagram

class Testanagram(unittest.TestCase):
    def test_apply_anagram(self):
        """
        tesh if the anagram output is matched with the expected output
        """
        self.assertEqual(
        apply_anagram(['tacio', 'tac', 'act', 'cat', 'catx', \
            'xcat', 'atxc', 'aimn', 'iamn', 'main', 'xcot']), 
        {'aciot': ['tacio'], 'act': ['tac', 'act', 'cat'], \
            'actx': ['catx', 'xcat', 'atxc'], 'aimn': ['aimn', 'iamn', 'main'], 'cotx': ['xcot']}, 
        "output is not matched")
    
if __name__ == '__main__':
    unittest.main()
