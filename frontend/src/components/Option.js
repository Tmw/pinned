import React from "react";
import PropTypes from "prop-types";
import Image from "react-graceful-image";

const Option = ({ avatar, name, onClick }) => (
  <button className="option" onClick={onClick}>
    <figure className="option--avatar">
      <Image
        src={avatar}
        alt={`avatar of ${name}`}
        width={64}
        height={64}
        placeholderColor="#DEDEDE"
      />
    </figure>

    <div className="option--name">
      <p>{name}</p>
    </div>

    <div className="option--arrow">
      <span>&rarr;</span>
    </div>
  </button>
);

Option.propTypes = {
  avatar: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired
};

export default Option;
