import React from "react";
import PropTypes from "prop-types";

const STAR_EMPTY = "☆";
const STAR_FILLED = "★";

const makeArray = (length, content) => Array(length).fill(content);

const makeStarsString = (stars, total) =>
  [
    ...makeArray(stars, STAR_FILLED),
    ...makeArray(total - stars, STAR_EMPTY)
  ].reduce((acc, slot) => acc + slot, "");

const StarBox = ({ stars, total }) => (
  <div className="score--stars">{makeStarsString(stars, total)}</div>
);

StarBox.propTypes = {
  stars: PropTypes.number.isRequired,
  total: PropTypes.number.isRequired
};

export default StarBox;
