import _ from 'lodash';
import { processTsWithLatencyBreakdown } from '../js/components/util/MetricUtils.js';
import latencyFixtures from './fixtures/latencyTs.json';

describe('MetricUtils', () => {
  describe('processTsWithLatencyBreakdown', () => {
    it('Converts raw metrics to plottable timeseries data', () => {
      let result = processTsWithLatencyBreakdown(latencyFixtures.metrics);
      let histograms = ['P50', 'P95', 'P99'];
      let expectedTsLength = _.size(histograms) * _.size(latencyFixtures.metrics[0].datapoints);

      expect(_.size(result["LATENCY"])).to.equal(expectedTsLength);
      _.each(result["LATENCY"], datum => {
        expect(datum.timestamp).not.to.be.empty;
        expect(datum.value).not.to.be.empty;
        expect(datum.label).to.be.oneOf(histograms);
      });
    });
  });
});
