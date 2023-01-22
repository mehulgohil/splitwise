import {Button, FormControl, Grid, TextField} from "@mui/material";
import {useState} from "react";

function CreateTransaction() {
  const [totalAmount, setTotalAmount] = useState();
  const [place, setPlace] = useState();
  const [date, setDate] = useState();
  const [spentBy, setSpentBy] = useState();
  const [spentByMobileNo, setSpentByMobileNo] = useState();
  const [lentTo, setLentTo] = useState();
  const [lentToMobileNo, setLentToMobileNo] = useState();
  const [shareAmt, setShareAmt] = useState();

  const onFieldValueChange = (e, fieldType) => {
    switch (fieldType) {
      case "TotalAmt":
        setTotalAmount(e.target.value)
        break
      case "Place":
        setPlace(e.target.value)
        break
      case "Date":
        setDate(e.target.value)
        break
      case "SpentBy":
        setSpentBy(e.target.value)
        break
      case "SpentByMobileNo":
        setSpentByMobileNo(e.target.value)
        break
      case "LentTo":
        setLentTo(e.target.value)
        break
      case "LentToMobileNo":
        setLentToMobileNo(e.target.value)
        break
      case "ShareAmt":
        setShareAmt(e.target.value)
        break
    }
  }

  async function submitPostTransaction() {
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        totalAmount: parseInt(totalAmount),
        place: place,
        date: date,
        spentBy: {
          mobile: spentByMobileNo,
          name: spentBy
        },
        nPeople: 1,
        split: [
          {
            mobile: lentToMobileNo,
            name: lentTo,
            shareAmount: parseInt(shareAmt)
          }
        ]
      })
    };
    const response = await fetch('http://localhost:8080/transactions', requestOptions);
    const data = await response.json();
    alert(data)
  }

  return (
    <FormControl style={{marginLeft: "20px"}}>
      <Grid container spacing={2}>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Total Amount" variant="standard" onChange={(e) => onFieldValueChange(e, "TotalAmt")}/>
        </Grid>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Place" variant="standard" onChange={(e) => onFieldValueChange(e, "Place")}/>
        </Grid>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Date" variant="standard" onChange={(e) => onFieldValueChange(e, "Date")}/>
        </Grid>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Spent By" variant="standard" onChange={(e) => onFieldValueChange(e, "SpentBy")}/>
        </Grid>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Mobile No" variant="standard" onChange={(e) => onFieldValueChange(e, "SpentByMobileNo")}/>
        </Grid>
        <Grid item xs={4}>
        </Grid>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Lent To" variant="standard" onChange={(e) => onFieldValueChange(e, "LentTo")}/>
        </Grid>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Mobile No" variant="standard" onChange={(e) => onFieldValueChange(e, "LentToMobileNo")}/>
        </Grid>
        <Grid item xs={4}>
          <TextField id="standard-basic" label="Share Amt" variant="standard" onChange={(e) => onFieldValueChange(e, "ShareAmt")}/>
        </Grid>
      </Grid>
      <Button variant="contained" style={{marginTop: "10px"}} onClick={submitPostTransaction}>Submit</Button>
    </FormControl>
  )
}

export default CreateTransaction