import {Button, FormControl, TextField} from "@mui/material";
import {useState} from "react";

function ChangePaymentStatus() {

  const [mobileNo, setMobileNo] = useState();
  const [transactionId, setTransactionId] = useState()

  function onTransactionIdChange(e) {
    console.log(e.target.value)
    setTransactionId(e.target.value)
  }

  function onMobileChange(e) {
    console.log(e.target.value)
    setMobileNo(e.target.value)
  }

  async function changeTransactionStatus() {
    const requestOptions = {
      method: 'PATCH',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ mobile: mobileNo })
    };
    const response = await fetch('http://localhost:8080/transactions/'+transactionId, requestOptions);
    const data = await response.json();
    alert(data)
  }

  return (
    <FormControl style={{marginLeft: "20px"}}>
      <TextField id="standard-basic" label="TransactionId" variant="standard" onChange={onTransactionIdChange}/>
      <TextField id="standard-basic" label="Mobile No" variant="standard" onChange={onMobileChange}/>
      <Button variant="contained" style={{marginTop: "10px"}} onClick={changeTransactionStatus}>Submit</Button>
    </FormControl>
  )

}

export default ChangePaymentStatus