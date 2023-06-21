
# extract_orgchart_data.py

How to run : 

    pip install -r requirements.txt
    python extract_orgchart_data.py



extract_orgchart_data.py can extract data (ministers and the corresponding departments) from the gazette PDFs which the data is in tabular format. The script downloads all the PDFs(English version) from that website directly and extracts data from those gazettes. 

The script works as follows: 
1. Download Gazette PDFs from the cabinetoffice website
2. Convert those to docx
3. Iterate through the tables and extract the Ministry and the departments(Which is in the Column II in tables)
4. Save the extracted data in scripts/orgchart CSV format.


Limitations : 

 - This script is a rule-based one. So, it can only extract data from the gazette PDFs which is in tabular format in the given structure.


Since this is a rule based script, it finds only the departments and ministers in the given structure. When the PDF is converted to docx, there may be some formatting issues in the tables (There is no ideal way to extract content from PDFs with exact same formatting). So, it will miss some data when extracting. Overall, it can extract all the ministries and most of the departments - there are some missing departments as well.

Improvements:


 - The gazettes in the  cabinetoffice website contains two types of data. 

       1. Data in tabular format - mentioning the ministry and relevant departments under it in a table. Ex: gazette 2022-07-22
       2. Data in Sentences - mentioning the amendments of the last gazette. (Like.. These departments should be removed from this ministry and these departments should be added under this). Ex: gazette 2022-10-05

 - From this script, it can only identify the data in tabular format. Need to address the gazettes which the data is in sentences as well.

 - When a PDF is converted into a docx, there are some miss-formattings in the tables. Because of this, the script can not identify "Column II" which contains the departments in some tables. Therefore, it can not extract some departments. This needs to be fixed.

NOTE : Check extracted data before populate the GIG db