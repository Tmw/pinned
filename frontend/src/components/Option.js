import React from "react";
import PropTypes from "prop-types";
import Image from "react-graceful-image";

const Option = ({ avatar, name, onClick }) => (
  <button className="button option" onClick={onClick}>
    <figure className="option--avatar">
      <Image
        src={avatar}
        alt={`avatar of ${name}`}
        width={64}
        height={64}
        style={{ objectFit: "cover", objectPosition: "Top" }}
        placeholderColor="#DEDEDE"
      />
    </figure>

    <div className="option--name">{name}</div>
    <div className="option--arrow">&rarr;</div>
  </button>
);

Option.propTypes = {
  avatar: PropTypes.string.isRequired,
  name: PropTypes.string.isRequired,
  onClick: PropTypes.func.isRequired
};

export default Option;
