import React from "react";
import PropTypes from "prop-types";

import WithSlackLinks from "components/WithSlackLinks";
import EmojiToUnicode from "components/EmojiToUnicode";

const PinBox = ({ text }) => (
  <div className="quote--box">
    <blockquote className="quote">
      <WithSlackLinks highlight>{EmojiToUnicode(text)}</WithSlackLinks>
    </blockquote>
  </div>
);

PinBox.propTypes = {
  text: PropTypes.string.isRequired
};

export default PinBox;
