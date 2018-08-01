import React from "react";
import { inject, observer } from "mobx-react";

class PinView extends React.Component {
  renderOptions() {
    const { ChallangeStore } = this.props.store;
    const { options } = ChallangeStore.currentChallange;

    return options.map((option, i) => (
      <li className="option" key={i} onClick={this.optionClicked}>
        <img
          className="option--avatar"
          src={option.avatar}
          alt={`avatar of ${option.name}`}
        />
        <p className="option--name">{option.name}</p>
      </li>
    ));
  }

  optionClicked = e => {
    this.props.store.ChallangeStore.nextChallange();
  };

  render() {
    const { ChallangeStore } = this.props.store;
    const challange = ChallangeStore.currentChallange;
    const challangeIndex = ChallangeStore.currentChallangeIdx + 1;
    const numChallanges = ChallangeStore.challanges.length;

    return (
      <div className="challangeView">
        <div className="progress--indicator">
          <span>
            {challangeIndex} / {numChallanges}
          </span>
        </div>

        <div className="quote--box">
          <blockquote className="quote">{challange.text}</blockquote>
        </div>

        <ul className="option--list">{this.renderOptions()}</ul>
      </div>
    );
  }
}

export default inject("store")(observer(PinView));
