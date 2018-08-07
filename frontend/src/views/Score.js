import { inject, observer } from "mobx-react";
import React from "react";

class ScoreView extends React.Component {
  render() {
    const { ChallangeStore } = this.props.store;

    const countScore = (score, challange) =>
      challange.answeredAuthorId === challange.author.id ? score + 1 : score;

    const totalScore = ChallangeStore.challanges.reduce(countScore, 0);

    return <h1>Your Score is: {totalScore}</h1>;
  }
}

export default inject("store")(observer(ScoreView));
