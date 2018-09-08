import React from "react";
import PropTypes from "prop-types";

const Spinner = ({ text }) => (
  <div className="spinner spinner--wrapper">
    <div className="spinner spinner--disc" />
    <p className="spinner spinner--text">{text}</p>
  </div>
);

Spinner.propTypes = {
  text: PropTypes.string.isRequired
};

export default Spinner;
