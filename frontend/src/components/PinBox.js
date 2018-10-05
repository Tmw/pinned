import React from "react";
import PropTypes from "prop-types";

import WithSlackLinks from "components/WithSlackLinks";
import EmojiToUnicode from "components/EmojiToUnicode";

const PinBox = ({ text }) => (
  <blockquote className="quote">
    <WithSlackLinks highlighted>{EmojiToUnicode(text)}</WithSlackLinks>
  </blockquote>
);

PinBox.propTypes = {
  text: PropTypes.string.isRequired
};

export default PinBox;
