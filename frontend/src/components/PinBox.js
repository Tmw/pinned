import React from "react";
import { emojify } from "react-emojione";
import PropTypes from "prop-types";

const PinBox = ({ text }) => (
  <div className="quote--box">
    <blockquote className="quote">
      {emojify(text, { output: "unicode" })}
    </blockquote>
  </div>
);

PinBox.propTypes = {
  text: PropTypes.string.isRequired
};

export default PinBox;
