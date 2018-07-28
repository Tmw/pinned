import React from "react";

export default ({ text }) => (
  <div className="spinner spinner--wrapper">
    <div className="spinner spinner--disc" />
    <p className="spinner spinner--text">{text}</p>
  </div>
);
