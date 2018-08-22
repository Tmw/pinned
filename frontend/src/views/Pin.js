import React from "react";
import { inject, observer } from "mobx-react";

class PinView extends React.Component {
  renderOptions() {
    const { ChallengeStore } = this.props.store;
    const { options } = ChallengeStore.currentChallenge;

    return options.map((option, i) => (
      <li
        className="option"
        key={i}
        onClick={e => this.optionClicked(option.id)}
      >
        <img
          className="option--avatar"
          src={option.avatar}
          alt={`avatar of ${option.name}`}
        />
        <p className="option--name">{option.name}</p>
      </li>
    ));
  }

  optionClicked = guessId => {
    this.props.store.ChallengeStore.currentChallenge.answer(guessId);
  };

  render() {
    const { ChallengeStore } = this.props.store;
    const challenge = ChallengeStore.currentChallenge;
    const challengeIndex = ChallengeStore.currentChallengeIndex + 1;
    const numChallenges = ChallengeStore.challenges.length;

    return (
      <div className="challenge-view">
        <div className="progress--indicator">
          <span>
            {challengeIndex} / {numChallenges}
          </span>
        </div>

        <div className="quote--box">
          <blockquote className="quote">{challenge.text}</blockquote>
        </div>

        <ul className="option--list">{this.renderOptions()}</ul>
      </div>
    );
  }
}

export default inject("store")(observer(PinView));
