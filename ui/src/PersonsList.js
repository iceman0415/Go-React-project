import React from "react";
import Person from "./Person";

function PersonsList(props) {
  return (
    <div>{props.persons.map(item => <Person key={item.email} name={item.name} age={item.age} balance={item.balance} email={item.email} address={item.address} checked={item.checked}/>)}</div> 
  ); 
}

export default PersonsList;