# 2006swe60
SC2006 Software Engineering Group 60

- Ng Wei Yu
- Park Yumin 
- Tay Jih How
- Saravanan Deepika 
- See Tow Tze Jiet 
- Shen Jia Cheng

Setting up Localhost
---------------------
Things to download before proceeding:
- [WampServer](https://www.wampserver.com/en/#download-wrapper) (includes PHP,Apache and MySQL)

After downloading and installing WampServer, run it wait till you see the green icon which means its good to go

![image](https://github.com/user-attachments/assets/41e47064-c30e-422d-bd77-aff87afa23bf)

Navigate to the directory of the wampserver which will be under
>C:\wamp64\www

This folder will be the place where the localhost will load the website. <br />
Open the browser of your choice and enter localhost in the url tab which you will see this page
![image](https://github.com/user-attachments/assets/a42ead09-89d4-428e-91d2-c5bae5d43e66)
if you see this page it means you are good to go

Now you can either clone the repository or download the whole folder and paste it under the www folder
![image](https://github.com/user-attachments/assets/d5790d5e-d0b7-45e5-9190-a017a6be6daf)

So in order to access the contents that you just pasted into the folder just type this into the url bar e.g sc2006 folder will be
>localhost/sc2006

After that you will see the website like this
![image](https://github.com/user-attachments/assets/59acf2e2-580a-4269-86db-40927dcce602)
Lastly we have to set up the database to store the user so we will go to PhpMyAdmin by using the url bar
>localhost/phpmyadmin

![image](https://github.com/user-attachments/assets/ccfc1011-68f7-42c6-9899-54dfd0d9fe75)
after you see this for now the account by default are <br />
username: root <br />
password: (blank) <br />
Server Choice: MySQL <br />
Click Login after that and you will come to this page. <br />
Next you have to import the .sql file inside the folder that you downloaded
![image](https://github.com/user-attachments/assets/fde0053f-264c-42ea-b694-ad81e1db9d40)
Choose the recyclo.sql file and scroll down to import the database
![image](https://github.com/user-attachments/assets/28bdab2c-cacb-4c9a-8c05-5dc0e0b969b5)
If successful you will see the recyclo database on the left sidebar
![image](https://github.com/user-attachments/assets/788ea105-2737-490d-9f12-e4c709c7925f)
Inside contains an account which was already created for the user test
