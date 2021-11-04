# nosqli-login-cracker
This go script has been designed after a TryHackMe challenge : https://tryhackme.com/room/nosqlinjectiontutorial
It is designed to crack the login page supplied in this particular challenge through NoSQL injection.
However, the script template can be re-utilized and adapted to match any kind of NoSQL injection. 

No interactive mode is implemented, so feel free to mess around with the variables as you please.

# How it works
The script begins by searching for the password length through the following payload : <SNIP>&pass[$regex]=^.{i}$</SNIP> .
Once the length is determined, it cracks the password by iterating through all possible characters and validating each possibility via <SNIP>&pass[$regex]=^abc......$</SNIP> .
Once the cracking is done, the password gets displayed to the screen.

# License 
No fucks given do whatever you want
