import React from "react";
import PropTypes from "prop-types";

const Option = ({ avatar, name, onClick }) => (
  <li className="option" onClick={onClick}>
    <img className="option--avatar" src={avatar} alt={`avatar of ${name}`} />
    <p className="option--name">{name}</p>
  </li>
);

Option.propTypes = {
  avatar: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired
};

export default Option;
