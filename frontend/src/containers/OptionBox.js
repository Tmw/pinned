import React from "react";
import Option from "components/Option";

class OptionBox extends React.Component {
  renderOptions() {
    return this.props.options.map((choice, i) => (
      <Option
        key={i}
        id={choice.id}
        onClick={() => this.props.onOptionClicked(choice.id)}
        avatar={choice.avatar}
        name={choice.name}
      />
    ));
  }

  render() {
    return <div className="option--list">{this.renderOptions()}</div>;
  }
}

export default OptionBox;
