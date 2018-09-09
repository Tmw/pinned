import React from "react";
import PropTypes from "prop-types";
import Image from "react-graceful-image";

const Option = ({ avatar, name, onClick }) => (
  <li className="option" onClick={onClick}>
    <Image
      src={avatar}
      alt={`avatar of ${name}`}
      width={192}
      height={192}
      className="option--avatar"
    />
    <p className="option--name">{name}</p>
  </li>
);

Option.propTypes = {
  avatar: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired
};

export default Option;
