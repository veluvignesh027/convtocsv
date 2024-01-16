# convtocsv
Tool to convert the trivy report into ".csv" file.


_Steps to convert TRIVY report into csv:_


**1. Clone the repo using the URL:**
  git clone https://github.com/veluvignesh027/convtocsv.git

**2. Run the below commands:**

  sed -i -e 's/\r$//' conv2csv
  
  chmod +x conv2csv | cp conv2csv /usr/bin


**3. Run the script:**
  conv2csv <imagename>

**4. Edit the temp.json file to the format mentioned below:**


**5. Run go:**  
  go run .
