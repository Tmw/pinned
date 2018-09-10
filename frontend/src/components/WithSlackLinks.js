import React from "react";
import PropTypes from "prop-types";

const SLACK_LINK_REGEX = /<([#@:/a-zA-Z0-9.]+)+>/g;
const toHighlighted = text => `<span class='highlight'>${text}</span>`;

const WithSlackLinks = ({ children, highlight = false }) => {
  const parsed = children.replace(
    SLACK_LINK_REGEX,
    (_, ...groups) => (highlight ? toHighlighted(groups[0]) : groups[0])
  );

  return <span dangerouslySetInnerHTML={{ __html: parsed }} />;
};

WithSlackLinks.propTypes = {
  children: PropTypes.string.isRequired,
  highlight: PropTypes.bool
};

export default WithSlackLinks;
