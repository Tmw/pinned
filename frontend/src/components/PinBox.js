import React from "react";
import { emojify } from "react-emojione";
import PropTypes from "prop-types";

const highlightRegex = /<([#@:/a-zA-Z0-9.]+)+>/g;
const HighLightLinks = text =>
  text.replace(
    highlightRegex,
    (full, ...groups) => `<span class='highlight'>${groups[0]}</span>`
  );

const PinBox = ({ text }) => (
  <div className="quote--box">
    <blockquote
      className="quote"
      dangerouslySetInnerHTML={{
        __html: HighLightLinks(emojify(text, { output: "unicode" }))
      }}
    />
  </div>
);

PinBox.propTypes = {
  text: PropTypes.string.isRequired
};

export default PinBox;
