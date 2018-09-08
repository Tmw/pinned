import React from "react";
import { inject, observer } from "mobx-react";

import ProgressIndicator from "components/ProgressIndicator";
import PinBox from "components/PinBox";
import OptionBox from "containers/OptionBox";
import PageTitle from "components/PageTitle";

class PinView extends React.Component {
  get challengeStore() {
    return this.props.store.ChallengeStore;
  }

  optionClickedHandler = optionID => {
    this.challengeStore.currentChallenge.answer(optionID);
  };

  render() {
    const challenge = this.challengeStore.currentChallenge,
      challengeIndex = this.challengeStore.currentChallengeIndex + 1,
      numChallenges = this.challengeStore.challenges.length;

    return (
      <React.Fragment>
        <PageTitle subtitle={`pin ${challengeIndex} / ${numChallenges}`} />

        <div className="challenge-view">
          <ProgressIndicator step={challengeIndex} total={numChallenges} />
          <PinBox text={challenge.text} />
          <OptionBox
            options={challenge.choices}
            onOptionClicked={this.optionClickedHandler}
          />
        </div>
      </React.Fragment>
    );
  }
}

export default inject("store")(observer(PinView));
