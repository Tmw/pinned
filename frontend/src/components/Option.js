import React from "react";
import PropTypes from "prop-types";
import Image from "react-graceful-image";

const Option = ({ avatar, name, onClick }) => (
  <div className="option" onClick={onClick}>
    <figure className="option--avatar">
      <Image
        src={avatar}
        alt={`avatar of ${name}`}
        width={64}
        height={64}
        placeholderColor="#DEDEDE"
      />
    </figure>
    <p className="option--name">{name}</p>
  </div>
);

Option.propTypes = {
  avatar: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired
};

export default Option;
