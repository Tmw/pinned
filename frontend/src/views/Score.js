import { inject, observer } from "mobx-react";
import React from "react";

class ScoreView extends React.Component {
  render() {
    const { ChallengeStore } = this.props.store;

    const countScore = (score, challenge) =>
      challenge.isCorrect ? score + 1 : score;

    const totalScore = ChallengeStore.challenges.reduce(countScore, 0);

    return <h1>Your Score is: {totalScore}</h1>;
  }
}

export default inject("store")(observer(ScoreView));
