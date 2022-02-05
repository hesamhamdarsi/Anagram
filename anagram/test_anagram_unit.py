import unittest
from anagrampkg.anagramlib import apply_anagram

class Testanagram(unittest.TestCase):
    def test_apply_anagram(self):
        """
        tesh the anagram output 
        """
        self.assertEqual(
        apply_anagram(['tacio', 'tac', 'act', 'cat', 'catx', 'xcat', 'atxc', 'aimn', 'iamn', 'main', 'xcot']), 
        [['tacio'], ['tac', 'act', 'cat'], ['catx', 'xcat', 'atxc'], ['aimn', 'iamn', 'main'], ['xcot']], 
        "output is not matched")
    
if __name__ == '__main__':
    unittest.main()
