import {Button, FormControl, TextField} from "@mui/material";
import {useState} from "react";

function GetOweBy() {
  const [mobileNo, setMobileNo] = useState();

  function onMobileChange(e) {
    setMobileNo(e.target.value)
  }

  async function getOweByDetails() {
    const requestOptions = {
      method: 'GET',
    };
    const response = await fetch('http://localhost:8080/transactions/oweBy/'+mobileNo, requestOptions);
    const data = await response.json();
    alert(JSON.stringify(data))
  }

  return (
    <FormControl style={{marginLeft: "20px"}}>
      <TextField id="standard-basic" label="Mobile No" variant="standard" onChange={onMobileChange}/>
      <Button variant="contained" style={{marginTop: "10px"}} onClick={getOweByDetails} >Submit</Button>
    </FormControl>
  )
}

export default GetOweBy