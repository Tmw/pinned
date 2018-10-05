import React from "react";
import PropTypes from "prop-types";

const ProgressIndicator = ({ step, total }) => (
  <div className="progress--indicator">
    {step} / {total}
  </div>
);

ProgressIndicator.propTypes = {
  step: PropTypes.number.isRequired,
  total: PropTypes.number.isRequired
};

export default ProgressIndicator;
