/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */

package pow

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_Reward(t *testing.T) {
	assert.Equal(t, GetReward(0), rewardTable[0])
	assert.Equal(t, GetReward(blockNumberPerEra-1), rewardTable[0])

	assert.Equal(t, GetReward(blockNumberPerEra), rewardTable[1])
	assert.Equal(t, GetReward(blockNumberPerEra*2-1), rewardTable[1])

	assert.Equal(t, GetReward(blockNumberPerEra*uint64(len(rewardTable))-1), rewardTable[len(rewardTable)-1])

	assert.Equal(t, GetReward(blockNumberPerEra*uint64(len(rewardTable))), tailReward)
}
