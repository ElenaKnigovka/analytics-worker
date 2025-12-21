const logger = require('./logger');
const { validateEvent, parseEventData } = require('./event-validator');

class Parser {
  constructor(config) {
    this.config = config;
  }

  async parseEvent(event) {
    try {
      const isValid = validateEvent(event);
      if (!isValid) {
        logger.error('Invalid event format');
        return null;
      }
      const eventData = parseEventData(event);
      return eventData;
    } catch (error) {
      logger.error('Error parsing event:', error);
      return null;
    }
  }
}

module.exports = Parser;