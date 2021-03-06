import React from "react";
import Option from "components/Option";

class OptionBox extends React.Component {
  renderOptions() {
    return this.props.options.map(choice => (
      <Option
        key={choice.id}
        id={choice.id}
        onClick={() => this.props.onOptionClicked(choice.id)}
        avatar={choice.avatar}
        name={choice.name}
      />
    ));
  }

  render() {
    return <div className="options">{this.renderOptions()}</div>;
  }
}

export default OptionBox;
