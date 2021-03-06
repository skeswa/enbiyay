import { SplashDataBuilder } from '../../../courtblink/api/builders/SplashDataBuilder'
import { ScoreboardCache } from '../../../nba/caches/ScoreboardCache'
import { TeamNewsCache } from '../../../nba/caches/TeamNewsCache'
import { Logger } from '../../../util/Logger'

import { CachedApiService } from './impl-cached'
import { ApiService, ApiServiceCreationStrategy } from './types'

/**
 * Creates a new API service.
 * @param logger logging utility.
 * @param scoreboardCache caches the NBA scoreboard.
 * @param splashDataBuilder constructs splash data.
 * @param teamNewsCache caches team news.
 * @return a new API service.
 */
export function createApiService(
  strategy: ApiServiceCreationStrategy.UsingCaches,
  logger: Logger,
  scoreboardCache: ScoreboardCache,
  splashDataBuilder: SplashDataBuilder,
  teamNewsCache: TeamNewsCache
): ApiService {
  return new CachedApiService(
    logger,
    scoreboardCache,
    splashDataBuilder,
    teamNewsCache
  )
}
