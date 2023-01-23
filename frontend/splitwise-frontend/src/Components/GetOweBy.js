import {Button, FormControl, TextField} from "@mui/material";
import {useState} from "react";

function GetOweBy() {
  const [mobileNo, setMobileNo] = useState();
  const [output, setOutput] = useState("")

  function onMobileChange(e) {
    setMobileNo(e.target.value)
  }

  async function getOweByDetails() {
    const requestOptions = {
      method: 'GET',
    };
    const response = await fetch('http://localhost:8080/transactions/oweBy/'+mobileNo, requestOptions);
    const data = await response.json();
    setOutput(JSON.stringify(data))
  }

  return (
    <FormControl style={{marginLeft: "20px", width: "50%"}}>
      <TextField id="standard-basic" label="Mobile No" variant="standard" onChange={onMobileChange}/>
      <Button variant="contained" style={{marginTop: "10px"}} onClick={getOweByDetails} >Submit</Button>
      <TextField id="standard-basic" label="output" variant="standard" value={output} multiline/>
    </FormControl>
  )
}

export default GetOweBy