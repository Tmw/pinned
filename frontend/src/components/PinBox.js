import React from "react";
import { emojify } from "react-emojione";

const PinBox = ({ text }) => (
  <div className="quote--box">
    <blockquote className="quote">
      {emojify(text, { output: "unicode" })}
    </blockquote>
  </div>
);

export default PinBox;
