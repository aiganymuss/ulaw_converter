#include <sox.h>
#include <stdio.h>
#include <malloc.h>

int decode(char* fileBuffer, size_t bufferSize, char* outputFileName) {
    // Set up SoX initialization
    sox_format_t* input = sox_open_mem_read(fileBuffer, bufferSize, NULL, NULL, "ul");

    if (input == NULL) {
        fprintf(stderr, "Error opening audio from memory buffer\n");
        free(fileBuffer);
        return 1;
    }

    // Set up the buffer to read audio data
    sox_sample_t* buffer = (sox_sample_t*)malloc(bufferSize * sizeof(sox_sample_t));
    if (buffer == NULL) {
        fprintf(stderr, "Error allocating buffer\n");
        sox_close(input);
        free(fileBuffer);
        return 1;
    }

    // Set up SoX output initialization
    sox_format_t* output = sox_open_write(outputFileName, &input->signal, NULL, "wav", NULL, NULL);

    if (output == NULL) {
        fprintf(stderr, "Error opening output audio file: %s\n", outputFileName);
        sox_close(input);
        free(fileBuffer);
        return 1;
    }

    // Read and write audio data
    size_t samples_read;
    while ((samples_read = sox_read(input, buffer, bufferSize / sizeof(sox_sample_t))) > 0) {
        sox_write(output, buffer, samples_read);
    }

    // Clean up
    free(fileBuffer);
    free(buffer);
    sox_close(input);
    sox_close(output);

    printf("Decoded audio data written to %s\n", outputFileName);

    return 0;
}
