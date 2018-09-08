import { types, flow, getRoot } from "mobx-state-tree";
import { Views } from "Constants";
import API from "services/API";

import Challenge from "state/Challenge";

const sumScore = (score, challenge) =>
  challenge.isCorrect ? score + 1 : score;

const ChallengeStore = types
  .model({
    challenges: types.array(Challenge)
  })

  .views(self => ({
    get currentChallengeIndex() {
      return self.challenges.filter(challenge => challenge.isAnswered).length;
    },

    get currentChallenge() {
      return self.challenges.find(challenge => !challenge.isAnswered);
    },

    get score() {
      return self.challenges.reduce(sumScore, 0);
    }
  }))

  .actions(self => ({
    fetchChallenges: flow(function*() {
      try {
        self.challenges = yield API.FetchChallenges();
        getRoot(self).ViewStore.presentView(Views.PIN);
      } catch (err) {
        getRoot(self).ViewStore.presentView(Views.ERROR, err.message);
      }
    })
  }));

const Defaults = {
  challenges: []
};

const CreateChallengeStore = overrides =>
  ChallengeStore.create(Object.assign(Defaults, overrides));

export { ChallengeStore, CreateChallengeStore };
