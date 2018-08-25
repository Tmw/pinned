import React from "react";
import { inject, observer } from "mobx-react";
import { emojify } from "react-emojione";

class PinView extends React.Component {
  renderChoices() {
    const { ChallengeStore } = this.props.store;
    const { choices } = ChallengeStore.currentChallenge;

    return choices.map((choice, i) => (
      <li
        className="option"
        key={i}
        onClick={e => this.optionClicked(choice.id)}
      >
        <img
          className="option--avatar"
          src={choice.avatar}
          alt={`avatar of ${choice.name}`}
        />
        <p className="option--name">{choice.name}</p>
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
          <blockquote className="quote">
            {emojify(challenge.text, { output: "unicode" })}
          </blockquote>
        </div>

        <ul className="option--list">{this.renderChoices()}</ul>
      </div>
    );
  }
}

export default inject("store")(observer(PinView));
