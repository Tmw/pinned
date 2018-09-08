import React from "react";

const ProgressIndicator = ({ step, total }) => (
  <div className="progress--indicator">
    <span>
      {step} / {total}
    </span>
  </div>
);

export default ProgressIndicator;
