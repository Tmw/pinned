import React from "react";

const Option = ({ id, avatar, name, onClick }) => (
  <li className="option" onClick={onClick}>
    <img className="option--avatar" src={avatar} alt={`avatar of ${name}`} />
    <p className="option--name">{name}</p>
  </li>
);
export default Option;
