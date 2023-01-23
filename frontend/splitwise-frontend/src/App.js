import * as React from 'react';
import {Box, Tab, Tabs} from "@mui/material";
import ChangePaymentStatus from "./Components/ChangePaymentStatus";
import TabPanel from "./Components/TabPanel";
import CreateTransaction from "./Components/CreateTransaction";
import GetOweBy from "./Components/GetOweBy";
import GetOweTo from "./Components/GetOweTo";
import Header from "./Components/Header";

function a11yProps(index) {
  return {
    id: `simple-tab-${index}`,
    'aria-controls': `simple-tabpanel-${index}`,
  };
}

function App() {

  const [value, setValue] = React.useState(0);

  const handleChange = (event, newValue) => {
    setValue(newValue);
  };

  return (
    <>
      <Header />
      <Box sx={{ width: '100%', typography: 'body1' }}>
        <Tabs value={value} onChange={(e, newVal) => handleChange(e, newVal)} aria-label="basic tabs example">
          <Tab label="Create Transaction" {...a11yProps(0)}/>
          <Tab label="Change Status" {...a11yProps(1)}/>
          <Tab label="Get Owe By" {...a11yProps(2)}/>
          <Tab label="Get Owe To" {...a11yProps(3)}/>
        </Tabs>
        <TabPanel value={value} index={0}>
          <CreateTransaction />
        </TabPanel>
        <TabPanel value={value} index={1}>
          <ChangePaymentStatus />
        </TabPanel>
        <TabPanel value={value} index={2}>
          <GetOweBy />
        </TabPanel>
        <TabPanel value={value} index={3}>
          <GetOweTo />
        </TabPanel>
      </Box>
    </>
  );
}

export default App;
