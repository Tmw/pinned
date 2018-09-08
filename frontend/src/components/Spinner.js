import React from "react";

const Spinner = ({ text }) => (
  <div className="spinner spinner--wrapper">
    <div className="spinner spinner--disc" />
    <p className="spinner spinner--text">{text}</p>
  </div>
);

export default Spinner;
